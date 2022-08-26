package redis

const (
	Prefix = "blog:"
	KeyPostTimeZSet = "post:time"
	KeyPostScoreZSet = "post:score"
	KeyPostVotedZSetPF = "post:voted:"

	KeyCommunitySetPF = "community:"
)

func getRedisKey(key string) string {
	return Prefix + key
}