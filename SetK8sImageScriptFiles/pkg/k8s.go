package kube

import (
	"fmt"

	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func SetYamlImageTag(k8sYamlObj *yaml.RNode, image *string) (string, error) {

	containers, err := k8sYamlObj.Pipe(yaml.Lookup("spec", "template", "spec", "containers"))

	if err != nil {
		return "", err
	}

	err = containers.VisitElements(func(k8sYamlObj *yaml.RNode) error {
		err := k8sYamlObj.PipeE(yaml.SetField("image", yaml.NewScalarRNode(*image)))

		if err != nil {
			return fmt.Errorf("failed to set image field: %w", err)
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	modifiedK8sYaml, err := k8sYamlObj.String()

	if err != nil {
		return "", err
	}

	return modifiedK8sYaml, nil
}
