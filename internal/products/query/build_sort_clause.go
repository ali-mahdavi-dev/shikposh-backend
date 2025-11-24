package query

// buildSortClause builds Elasticsearch sort clause
func (h *ProductQueryHandler) buildSortClause(sort string) []map[string]interface{} {
	switch sort {
	case "price_asc":
		return []map[string]interface{}{
			{"price": map[string]interface{}{"order": "asc"}},
		}
	case "price_desc":
		return []map[string]interface{}{
			{"price": map[string]interface{}{"order": "desc"}},
		}
	case "rating":
		return []map[string]interface{}{
			{"rating": map[string]interface{}{"order": "desc"}},
		}
	case "newest":
		return []map[string]interface{}{
			{"created_at": map[string]interface{}{"order": "desc"}},
		}
	default:
		// Default: relevance score (no sort clause needed)
		return []map[string]interface{}{}
	}
}

