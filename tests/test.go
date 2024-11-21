package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Estructura para parsear la respuesta de Overpass para la ubicación
type OverpassLocationResponse struct {
	Elements []struct {
		Type  string `json:"type"`
		ID    int    `json:"id"`
		Tags  struct {
			Name       string `json:"name"`
			AdminLevel string `json:"admin_level,omitempty"`
			Boundary   string `json:"boundary,omitempty"`
		} `json:"tags"`
		Bounds struct {
			MinLat float64 `json:"minlat"`
			MinLon float64 `json:"minlon"`
			MaxLat float64 `json:"maxlat"`
			MaxLon float64 `json:"maxlon"`
		} `json:"bounds"`
	} `json:"elements"`
}

// Estructura para parsear la respuesta de Overpass para atracciones turísticas
type OverpassAttractionsResponse struct {
	Elements []struct {
		Type    string  `json:"type"`
		ID      int     `json:"id"`
		Lat     float64 `json:"lat,omitempty"`
		Lon     float64 `json:"lon,omitempty"`
		Center  struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"center,omitempty"`
		Tags struct {
			Name string `json:"name,omitempty"`
		} `json:"tags"`
	} `json:"elements"`
}

func main() {
	locationName := "Paris" // Cambia esto según la ubicación deseada

	// Paso 1: Obtener la relación administrativa de la ubicación
	locationURL := buildOverpassURL(`
		[out:json];
		relation["boundary"="administrative"]["name"="` + locationName + `"];
		out body;
		>;
		out skel qt;
	`)

	locationResp, err := fetchOverpassData(locationURL)
	if err != nil {
		log.Fatalf("Error fetching location data: %v", err)
	}

	location, err := parseLocationResponse(locationResp)
	if err != nil {
		log.Fatalf("Error parsing location response: %v", err)
	}

	if len(location.Elements) == 0 {
		log.Fatalf("No se encontró la ubicación: %s", locationName)
	}

	// Imprimir detalles de la ubicación
	fmt.Printf("Ubicación: %s\n", location.Elements[0].Tags.Name)
	fmt.Printf("Bounding Box:\n")
	fmt.Printf("  MinLat: %f, MinLon: %f\n", location.Elements[0].Bounds.MinLat, location.Elements[0].Bounds.MinLon)
	fmt.Printf("  MaxLat: %f, MaxLon: %f\n", location.Elements[0].Bounds.MaxLat, location.Elements[0].Bounds.MaxLon)

	// Paso 2: Si la ubicación es una ciudad (admin_level >= 8), obtener el país padre
	var countryName string
	if location.Elements[0].Tags.AdminLevel == "8" || location.Elements[0].Tags.AdminLevel == "6" {
		// admin_level=8 generalmente corresponde a ciudades
		// admin_level=6 puede variar según el país
		// Necesitamos buscar la relación padre con admin_level=2
		countryURL := buildOverpassURL(`
			[out:json];
			(
				rel(around:0)["boundary"="administrative"]["admin_level"="2"];
			);
			out body;
		`)

		countryResp, err := fetchOverpassData(countryURL)
		if err != nil {
			log.Fatalf("Error fetching country data: %v", err)
		}

		country, err := parseLocationResponse(countryResp)
		if err != nil {
			log.Fatalf("Error parsing country response: %v", err)
		}

		for _, elem := range country.Elements {
			countryName = elem.Tags.Name
			break
		}

		if countryName == "" {
			log.Println("No se encontró el país correspondiente.")
		} else {
			fmt.Printf("País: %s\n", countryName)
		}
	} else if location.Elements[0].Tags.AdminLevel == "2" {
		// Si la ubicación ya es un país
		countryName = location.Elements[0].Tags.Name
		fmt.Printf("País: %s\n", countryName)
	}

	// Paso 3: Obtener el bounding box
	minLat := location.Elements[0].Bounds.MinLat
	minLon := location.Elements[0].Bounds.MinLon
	maxLat := location.Elements[0].Bounds.MaxLat
	maxLon := location.Elements[0].Bounds.MaxLon

	// Paso 4: Consultar atracciones turísticas dentro del bbox
	attractionsURL := buildOverpassURL(`
		[out:json];
		(
			node["tourism"="attraction"](` + fmt.Sprintf("%f,%f,%f,%f", minLat, minLon, maxLat, maxLon) + `);
			way["tourism"="attraction"](` + fmt.Sprintf("%f,%f,%f,%f", minLat, minLon, maxLat, maxLon) + `);
			relation["tourism"="attraction"](` + fmt.Sprintf("%f,%f,%f,%f", minLat, minLon, maxLat, maxLon) + `);
		);
		out center;
	`)

	attractionsResp, err := fetchOverpassData(attractionsURL)
	if err != nil {
		log.Fatalf("Error fetching attractions data: %v", err)
	}

	attractions, err := parseAttractionsResponse(attractionsResp)
	if err != nil {
		log.Fatalf("Error parsing attractions response: %v", err)
	}

	// Imprimir atracciones turísticas
	fmt.Printf("\nAtracciones Turísticas en %s:\n", locationName)
	for _, attraction := range attractions.Elements {
		name := attraction.Tags.Name
		lat := attraction.Lat
		lon := attraction.Lon
		if attraction.Type == "way" || attraction.Type == "relation" {
			lat = attraction.Center.Lat
			lon = attraction.Center.Lon
		}
		if name != "" {
			fmt.Printf("- %s (Lat: %f, Lon: %f)\n", name, lat, lon)
		}
	}
}

// Función para construir la URL de Overpass con la consulta dada
func buildOverpassURL(query string) string {
	baseURL := "https://overpass-api.de/api/interpreter"
	encodedQuery := url.QueryEscape(query)
	return baseURL + "?data=" + encodedQuery
}

// Función para realizar la solicitud HTTP a Overpass y obtener la respuesta
func fetchOverpassData(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 60 * time.Second, // Aumentar el timeout si es necesario
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error realizando la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Leer el cuerpo para obtener detalles del error
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("Error en la respuesta HTTP: %s - %s", resp.Status, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error leyendo el cuerpo de la respuesta: %v", err)
	}

	return body, nil
}

// Función para parsear la respuesta de ubicación
func parseLocationResponse(data []byte) (*OverpassLocationResponse, error) {
	var response OverpassLocationResponse
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, fmt.Errorf("Error deserializando JSON: %v", err)
	}
	return &response, nil
}

// Función para parsear la respuesta de atracciones turísticas
func parseAttractionsResponse(data []byte) (*OverpassAttractionsResponse, error) {
	var response OverpassAttractionsResponse
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, fmt.Errorf("Error deserializando JSON: %v", err)
	}
	return &response, nil
}
