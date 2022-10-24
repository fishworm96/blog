package redis

const (
	Prefix = "blog:"
	KeyPostTimeZSet = "post:time"
	KeyPostScoreZSet = "post:score"
	KeyPostVotedZSetPF = "post:voted:"

	KeyCommunitySetPF = "community:"

	KeyEmailSetPF = "email:"
)

func getRedisKey(key string) string {
	return Prefix + key
}