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
	objects  map[T]T
	children map[T][]T
	parents  map[T][]T
}

func NewRelationator() Relationator {
	return Relationator{
		objects:  map[T]T{},
		children: map[T][]T{},
		parents:  map[T][]T{},
	}
}

type T interface{}

func mapKeys(m map[T]struct{}) []T {
	keys := make([]T, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func (r *Relationator) GetChildren(id T) []T {
	return r.children[id]
}

func (r *Relationator) GetAllChildren(id T) []T {
	childs := []T{}

	for _, child := range r.GetChildren(id) {
		r.recurChildren(&childs, child)
	}

	return childs
}

func (r *Relationator) recurChildren(childs *[]T, id T) {
	*childs = append(*childs, id)

	for _, child := range r.GetChildren(id) {
		r.recurChildren(childs, child)
	}
}

func (r *Relationator) GetAllUniqueChildren(id T) []T {
	childs := map[T]struct{}{}

	for _, child := range r.GetChildren(id) {
		r.recurUniqueChildren(childs, child)
	}

	return mapKeys(childs)
}

func (r *Relationator) recurUniqueChildren(childs map[T]struct{}, id T) {
	if _, ok := childs[id]; ok {
		return
	}
	childs[id] = struct{}{}

	for _, child := range r.GetChildren(id) {
		r.recurUniqueChildren(childs, child)
	}
}

func (r *Relationator) GetAllParents(id T) []T {
	parents := []T{}

	for _, parent := range r.GetParents(id) {
		r.recurParents(&parents, parent)
	}

	return parents
}

func (r *Relationator) recurParents(parents *[]T, id T) {
	*parents = append(*parents, id)

	for _, parent := range r.GetParents(id) {
		r.recurParents(parents, parent)
	}
}

func (r *Relationator) GetAllUniqueParents(id T) []T {
	parents := map[T]struct{}{}

	for _, parent := range r.GetParents(id) {
		r.recurUniqueParents(parents, parent)
	}

	return mapKeys(parents)
}

func (r *Relationator) recurUniqueParents(parents map[T]struct{}, id T) {
	if _, ok := parents[id]; ok {
		return
	}
	parents[id] = struct{}{}

	for _, parent := range r.GetParents(id) {
		r.recurUniqueParents(parents, parent)
	}
}

func (r *Relationator) GetParents(id T) []T {
	return r.parents[id]
}

func (r *Relationator) Set(id T, obj T) {
	r.objects[id] = obj
}

func (r *Relationator) Get(id T) T {
	return r.objects[id]
}

func (r *Relationator) AddUniqueChild(id, child T) {
	r.addChild(id, child, true)
}

func (r *Relationator) AddChild(id, child T) {
	r.addChild(id, child, false)
}

func (r *Relationator) addChild(id, child T, unique bool) {
	if _, ok := r.objects[id]; !ok {
		r.Set(id, nil)
	}

	if _, ok := r.objects[child]; !ok {
		r.Set(child, nil)
	}

	if unique {
		if !IsInFace(r.children[id], child) {
			r.children[id] = append(r.children[id], child)
		}
	} else {
		r.children[id] = append(r.children[id], child)
	}

	if !IsInFace(r.parents[child], id) {
		r.parents[child] = append(r.parents[child], id)
	}
}

func (r *Relationator) AddUniqueParent(id, parent T) {
	r.addParent(id, parent, true)
}

func (r *Relationator) AddParent(id, parent T) {
	r.addParent(id, parent, false)
}

func (r *Relationator) addParent(id, parent T, unique bool) {
	if _, ok := r.objects[id]; !ok {
		r.Set(id, nil)
	}

	if _, ok := r.objects[parent]; !ok {
		r.Set(parent, nil)
	}

	if !unique {
		r.parents[id] = append(r.parents[id], parent)
	} else {
		if !IsInFace(r.parents[id], parent) {
			r.parents[id] = append(r.parents[id], parent)
		}
	}

	if !IsInFace(r.children[parent], id) {
		r.children[parent] = append(r.children[parent], id)
	}
}

func IsInFace(slice []T, val T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}

	return false
}

func (r Relationator) StringNode(id T) string {
	var sb strings.Builder
	key := id
	val := r.objects[id]
	sb.WriteString(fmt.Sprintf("%v: %+v\n", key, val))
	children := r.children[key]
	sb.WriteString(fmt.Sprintf("  childrn: %+v\n", children))
	parents := r.parents[key]
	sb.WriteString(fmt.Sprintf("  parents: %+v", parents))

	return sb.String()
}

func (r Relationator) String() string {
	var sb strings.Builder
	for key := range r.objects {
		if sb.Len() > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(r.StringNode(key))
	}
	return sb.String()
}
