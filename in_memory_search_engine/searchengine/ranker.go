package searchengine

import "sort"

type Ranker interface {
	rank(docs []*document, searchTerm string) []*document
}

type OccurenceRanker struct{}

func (o *OccurenceRanker) rank(docs []*document, searchTerm string) []*document {
	occurenceMap := make(map[*document]int)
	for _, doc := range docs {
		occurenceMap[doc] = doc.occurenceCount(searchTerm)
	}
	sort.Slice(docs, func(i, j int) bool {
		return occurenceMap[docs[i]] > occurenceMap[docs[j]]
	})
	return docs
}

type RecencyRanker struct{}

func (r *RecencyRanker) rank(docs []*document, searchTerm string) []*document {
	sort.Slice(docs, func(i, j int) bool {
		return docs[i].created.After(docs[j].created)
	})
	return docs
}
