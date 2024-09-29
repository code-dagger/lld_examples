package searchengine

type config struct {
	ranker   Ranker
	splitter Splitter
}

func newConfig(ranker Ranker, splitter Splitter) config {
	return config{
		ranker:   ranker,
		splitter: splitter,
	}
}
