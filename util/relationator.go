package util

import (
	"fmt"
	"strings"
)

// Relationator represents a set of relationships between objects, objects
// have parents and children
//
// When children are added, parent relationships are automatically set, and vice versa when
// parents are added
type Relationator struct {
	objects  map[string]interface{}
	children map[string][]string
	parents  map[string][]string
}

func NewRelationator() Relationator {
	return Relationator{
		objects:  map[string]interface{}{},
		children: map[string][]string{},
		parents:  map[string][]string{},
	}
}

func mapKeys(m map[string]struct{}) []string {
	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func (r *Relationator) GetChildren(id string) []string {
	return r.children[id]
}

func (r *Relationator) GetAllChildren(id string) []string {
	childs := []string{}

	for _, child := range r.GetChildren(id) {
		r.recurChildren(&childs, child)
	}

	return childs
}

func (r *Relationator) recurChildren(childs *[]string, id string) {
	*childs = append(*childs, id)

	for _, child := range r.GetChildren(id) {
		r.recurChildren(childs, child)
	}
}

func (r *Relationator) GetAllUniqueChildren(id string) []string {
	childs := map[string]struct{}{}

	for _, child := range r.GetChildren(id) {
		r.recurUniqueChildren(childs, child)
	}

	return mapKeys(childs)
}

func (r *Relationator) recurUniqueChildren(childs map[string]struct{}, id string) {
	if _, ok := childs[id]; ok {
		return
	}
	childs[id] = struct{}{}

	for _, child := range r.GetChildren(id) {
		r.recurUniqueChildren(childs, child)
	}
}

func (r *Relationator) GetAllParents(id string) []string {
	parents := []string{}

	for _, parent := range r.GetParents(id) {
		r.recurParents(&parents, parent)
	}

	return parents
}

func (r *Relationator) recurParents(parents *[]string, id string) {
	*parents = append(*parents, id)

	for _, parent := range r.GetParents(id) {
		r.recurParents(parents, parent)
	}
}

func (r *Relationator) GetAllUniqueParents(id string) []string {
	parents := map[string]struct{}{}

	for _, parent := range r.GetParents(id) {
		r.recurUniqueParents(parents, parent)
	}

	return mapKeys(parents)
}

func (r *Relationator) recurUniqueParents(parents map[string]struct{}, id string) {
	if _, ok := parents[id]; ok {
		return
	}
	parents[id] = struct{}{}

	for _, parent := range r.GetParents(id) {
		r.recurUniqueParents(parents, parent)
	}
}

func (r *Relationator) GetParents(id string) []string {
	return r.parents[id]
}

func (r *Relationator) Set(id string, obj interface{}) {
	r.objects[id] = obj
}

func (r *Relationator) AddUniqueChild(id, child string) {
	r.addChild(id, child, true)
}

func (r *Relationator) AddChild(id, child string) {
	r.addChild(id, child, false)
}

func (r *Relationator) addChild(id, child string, unique bool) {
	if _, ok := r.objects[id]; !ok {
		r.Set(id, nil)
	}

	if _, ok := r.objects[child]; !ok {
		r.Set(child, nil)
	}

	if unique {
		if !IsIn(r.children[id], child) {
			r.children[id] = append(r.children[id], child)
		}
	} else {
		r.children[id] = append(r.children[id], child)
	}

	if !IsIn(r.parents[child], id) {
		r.parents[child] = append(r.parents[child], id)
	}
}

func (r *Relationator) AddUniqueParent(id, parent string) {
	r.addParent(id, parent, true)
}

func (r *Relationator) AddParent(id, parent string) {
	r.addParent(id, parent, false)
}

func (r *Relationator) addParent(id, parent string, unique bool) {
	if _, ok := r.objects[id]; !ok {
		r.Set(id, nil)
	}

	if _, ok := r.objects[parent]; !ok {
		r.Set(parent, nil)
	}

	if !unique {
		r.parents[id] = append(r.parents[id], parent)
	} else {
		if !IsIn(r.parents[id], parent) {
			r.parents[id] = append(r.parents[id], parent)
		}
	}

	if !IsIn(r.children[parent], id) {
		r.children[parent] = append(r.children[parent], id)
	}
}

func (r Relationator) String() string {
	var sb strings.Builder
	for key, val := range r.objects {
		if sb.Len() > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprintf("%s: %+v\n", key, val))
		children := r.children[key]
		sb.WriteString(fmt.Sprintf("  childrn: %+v\n", children))
		parents := r.parents[key]
		sb.WriteString(fmt.Sprintf("  parents: %+v", parents))

	}
	return sb.String()
}
