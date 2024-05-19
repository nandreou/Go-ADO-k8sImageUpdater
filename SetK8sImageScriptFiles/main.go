package main

import (
	"flag"
	"fmt"
	kube "k8s/image/update/pkg"
	"log"
	"os"
	"strings"

	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func main() {

	/************* Set Input Flags *************/
	image := flag.String("image", "", "The new image tag to set in the YAML file")
	yamlFile := flag.String("file", "", "The path to the Kubernetes YAML file")

	flag.Parse()

	if *image == "" || *yamlFile == "" {
		log.Fatal("Both -image and -file arguments are required")
	}

	/************* Read Yaml File *************/
	yamlData, err := os.ReadFile(*yamlFile)

	if err != nil {
		log.Fatal(err)
	}

	/************* Setting Up Multi Document Yaml File *************/
	splitedYaml := strings.Split(string(yamlData), "---\n")

	if len(splitedYaml) > 1 {
		fmt.Println("More Than one Files In a Single Yaml")

		var yamlSlice []string

		for _, data := range splitedYaml {
			k8sYamlObj, err := yaml.Parse(data)
			if err != nil {
				log.Fatal(err)
			}

			kindNode := k8sYamlObj.Field("kind")

			switch kindNode.Value.YNode().Value {
			case "Deployment":
				modifiedK8sYaml, err := kube.SetYamlImageTag(k8sYamlObj, image)
				if err != nil {
					log.Fatal(err)
				}
				yamlSlice = append(yamlSlice, modifiedK8sYaml)
			default:
				yamlFile, err := k8sYamlObj.String()
				if err != nil {
					log.Fatal(err)
				}
				yamlSlice = append(yamlSlice, yamlFile)
			}
		}

		finalYaml := strings.Join(yamlSlice, "\n---\n")

		err = os.WriteFile(*yamlFile, []byte(finalYaml), 0644)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		/************* Setting Single Document Yaml File *************/
		k8sYamlObj, err := yaml.Parse(string(yamlData))
		if err != nil {
			log.Fatal(err)
		}

		modifiedK8sYaml, err := kube.SetYamlImageTag(k8sYamlObj, image)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(*yamlFile, []byte(modifiedK8sYaml), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
