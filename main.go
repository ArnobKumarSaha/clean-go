package main

import (
	"fmt"
	_ "github.com/Arnobkumarsaha/clean-go/go-basics"
	"math"
)

func main() {
	fmt.Println(math.MaxInt, " ", math.MaxInt32, " ", math.MaxInt64)

	//slice := []string{"a", "b", "c"}
	//var ns []*string
	//for idx, _ := range slice {
	//	ns = append(ns, &slice[idx])
	//}
	//
	//for _, s := range ns {
	//	fmt.Println(*s)
	//}
	pods := make(map[PodId]*BasicPodSpec)

	id1 := PodId{
		name:      "n1",
		namespace: "s1",
	}
	pods[id1] = &BasicPodSpec{
		Id: id1,
		Containers: []ContainerSpec{
			{
				image: "i11",
				resource: map[string]int64{
					"cpu": 11,
					"mem": 11,
				},
			},
			{
				image: "i12",
				resource: map[string]int64{
					"cpu": 12,
					"mem": 12,
				},
			},
		},
	}
	id2 := PodId{
		name:      "n2",
		namespace: "s2",
	}
	pods[id2] = &BasicPodSpec{
		Id: id2,
		Containers: []ContainerSpec{
			{
				image: "i21",
				resource: map[string]int64{
					"cpu": 21,
					"mem": 21,
				},
			},
			{
				image: "i22",
				resource: map[string]int64{
					"cpu": 22,
					"mem": 22,
				},
			},
		},
	}

	for _, pod := range pods {
		changePod(pod, "i12")
	}
	for _, pod := range pods {
		fmt.Printf("name = %v, namespace = %v => ", pod.Id.name, pod.Id.namespace)
		for _, container := range pod.Containers {
			fmt.Printf("[CONTAINER] image = %v, ", container.image)
			for s, i := range container.resource {
				fmt.Printf("resource %v %v, ", s, i)
			}
		}
		fmt.Println()
	}
}

func changePod(pod *BasicPodSpec, image string) {
	for _, con := range pod.Containers {
		if con.image != image {
			continue
		}
		for n, _ := range con.resource {
			if n == "cpu" {
				con.resource[n] = 111
			}
		}
	}
}

type PodId struct {
	name      string
	namespace string
}

type BasicPodSpec struct {
	Id         PodId
	Containers []ContainerSpec
}

type ContainerSpec struct {
	image    string
	resource Resources
}

type Resources map[string]int64
