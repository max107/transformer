package transformer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/falaleev/shared/transformer"
)

type Source struct {
	Name string
}

type Destination struct {
	FullName string
}

func transform(s Source) Destination                   { return Destination{FullName: s.Name} }
func transformPointerDst(s Source) *Destination        { return &Destination{FullName: s.Name} }
func transformPointerSrc(s *Source) Destination        { return Destination{FullName: s.Name} }
func transformPointerSrcAndDst(s *Source) *Destination { return &Destination{FullName: s.Name} }

func TestTransform(t *testing.T) {
	t.Parallel()

	t.Run("normal", func(t *testing.T) {
		t.Parallel()

		items := []Source{
			{Name: "one"},
			{Name: "two"},
			{Name: "three"},
		}
		result := transformer.List(items, transform)
		require.NotEmpty(t, result)
		for i, w := range items {
			require.Equal(t, w.Name, result[i].FullName)
		}
	})

	t.Run("pointer_src", func(t *testing.T) {
		t.Parallel()

		items := []*Source{
			{Name: "one"},
			{Name: "two"},
			{Name: "three"},
		}
		result := transformer.List(items, transformPointerSrc)
		require.NotEmpty(t, result)
		for i, w := range items {
			require.Equal(t, w.Name, result[i].FullName)
		}
	})

	t.Run("pointer_dst", func(t *testing.T) {
		t.Parallel()

		items := []Source{
			{Name: "one"},
			{Name: "two"},
			{Name: "three"},
		}
		result := transformer.List(items, transformPointerDst)
		require.NotEmpty(t, result)
		for i, w := range items {
			require.Equal(t, w.Name, result[i].FullName)
		}
	})

	t.Run("pointer_all", func(t *testing.T) {
		t.Parallel()

		items := []*Source{
			{Name: "one"},
			{Name: "two"},
			{Name: "three"},
		}
		result := transformer.List(items, transformPointerSrcAndDst)
		require.NotEmpty(t, result)
		for i, w := range items {
			require.Equal(t, w.Name, result[i].FullName)
		}
	})
}
