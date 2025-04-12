package generics

type HashSet[T comparable] map[T]struct{}

func NewHashSet[T comparable]() HashSet[T] {
	return make(HashSet[T])
}

func (h HashSet[T]) Add(item T) {
	h[item] = struct{}{}
}

func (h HashSet[T]) Remove(item T) {
	delete(h, item)
}

func (h HashSet[T]) Contains(item T) bool {
	_, exists := h[item]
	return exists
}

func (h HashSet[T]) Size() int {
	return len(h)
}

func (h HashSet[T]) Clear() {
	for key := range h {
		delete(h, key)
	}
}

func (h HashSet[T]) Keys() []T {
	keys := make([]T, len(h))
	i := 0
	for key := range h {
		keys[i] = key
		i++
	}
	return keys
}

func (h HashSet[T]) IsEmpty() bool {
	return len(h) == 0
}

func (h HashSet[T]) Union(other HashSet[T]) HashSet[T] {
	unionSet := NewHashSet[T]()
	for key := range h {
		unionSet.Add(key)
	}
	for key := range other {
		unionSet.Add(key)
	}
	return unionSet
}

func (h HashSet[T]) Intersection(other HashSet[T]) HashSet[T] {
	intersectionSet := NewHashSet[T]()
	for key := range h {
		if other.Contains(key) {
			intersectionSet.Add(key)
		}
	}
	return intersectionSet
}

func (h HashSet[T]) Difference(other HashSet[T]) HashSet[T] {
	differenceSet := NewHashSet[T]()
	for key := range h {
		if !other.Contains(key) {
			differenceSet.Add(key)
		}
	}
	return differenceSet
}

func (h HashSet[T]) IsSubset(other HashSet[T]) bool {
	for key := range h {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (h HashSet[T]) IsSuperset(other HashSet[T]) bool {
	for key := range other {
		if !h.Contains(key) {
			return false
		}
	}
	return true
}

func (h HashSet[T]) Equals(other HashSet[T]) bool {
	if h.Size() != other.Size() {
		return false
	}
	for key := range h {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
