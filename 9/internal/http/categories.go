package http

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	lru "github.com/hashicorp/golang-lru"
	"lecture-9/internal/models"
	"lecture-9/internal/store"
	"net/http"
)

type CategoryResource struct {
	store store.Store
	cache *lru.TwoQueueCache
}

func NewCategoryResource(store store.Store, cache *lru.TwoQueueCache) *CategoryResource {
	return &CategoryResource{
		store: store,
		cache: cache,
	}
}

func (cr *CategoryResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", cr.CreateCategory)
	r.Get("/", cr.AllCategories)

	return r
}

func (cr *CategoryResource) CreateCategory(w http.ResponseWriter, r *http.Request) {
	category := new(models.Category)
	if err := json.NewDecoder(r.Body).Decode(category); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := cr.store.Categories().Create(r.Context(), category); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	// Правильно пройтись по всем буквам и всем словам
	cr.cache.Purge() // в рамках учебного проекта полностью чистим кэш после создания новой категории

	w.WriteHeader(http.StatusCreated)
}

func (cr *CategoryResource) AllCategories(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	filter := &models.CategoriesFilter{}

	searchQuery := queryValues.Get("query")
	if searchQuery != "" {
		categoriesFromCache, ok := cr.cache.Get(searchQuery)
		if ok {
			render.JSON(w, r, categoriesFromCache)
			return
		}

		filter.Query = &searchQuery
	}

	categories, err := cr.store.Categories().All(r.Context(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	if searchQuery != "" {
		cr.cache.Add(searchQuery, categories)
	}
	render.JSON(w, r, categories)
}

//r.Post("/categories", )
//r.Get("/categories", )
//r.Get("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
//	idStr := chi.URLParam(r, "id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		fmt.Fprintf(w, "Unknown err: %v", err)
//		return
//	}
//
//	category, err := s.store.Categories().ByID(r.Context(), id)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		fmt.Fprintf(w, "DB err: %v", err)
//		return
//	}
//
//	render.JSON(w, r, category)
//})
//r.Put("/categories", func(w http.ResponseWriter, r *http.Request) {
//	category := new(models.Category)
//	if err := json.NewDecoder(r.Body).Decode(category); err != nil {
//		w.WriteHeader(http.StatusUnprocessableEntity)
//		fmt.Fprintf(w, "Unknown err: %v", err)
//		return
//	}
//
//	err := validation.ValidateStruct(
//		category,
//		validation.Field(&category.ID, validation.Required),
//		validation.Field(&category.Name, validation.Required),
//	)
//	if err != nil {
//		w.WriteHeader(http.StatusUnprocessableEntity)
//		fmt.Fprintf(w, "Unknown err: %v", err)
//		return
//	}
//
//	if err := s.store.Categories().Update(r.Context(), category); err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		fmt.Fprintf(w, "DB err: %v", err)
//		return
//	}
//})
//r.Delete("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
//	idStr := chi.URLParam(r, "id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		fmt.Fprintf(w, "Unknown err: %v", err)
//		return
//	}
//
//	if err := s.store.Categories().Delete(r.Context(), id); err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		fmt.Fprintf(w, "DB err: %v", err)
//		return
//	}
//})
