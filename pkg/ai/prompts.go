package ai

const (
	default_prompt = `简化以下由三重破折号分隔的 Kubernetes 错误消息，该消息以 --- %s --- 语言编写； --- %s ---。
	以不超过 280 个字符的分步方式提供最可能的解决方案。 按以下格式编写输出：
	错误：{在此解释错误}
	解决方案：{此处为逐步解决方案}
	`
	trivy_vuln_prompt = "Explain the following trivy scan result and the detail risk or root cause of the CVE ID, then provide a solution. Response in %s: %s"
	trivy_conf_prompt = "Explain the following trivy scan result and the detail risk or root cause of the security check, then provide a solution."
)

var PromptMap = map[string]string{
	"default":             default_prompt,
	"VulnerabilityReport": trivy_vuln_prompt, // for Trivy integration, the key should match `Result.Kind` in pkg/common/types.go
	"ConfigAuditReport":   trivy_conf_prompt,
}
