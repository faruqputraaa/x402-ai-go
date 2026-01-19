package ai

func Prompt(topic, tone, audience string) string {
	return `
Kamu adalah social media strategist profesional.

Tugas:
Buat 1 caption Threads dalam Bahasa Indonesia.

Topik: ` + topic + `
Gaya bahasa: ` + tone + `
Target audience: ` + audience + `

Aturan:
- Santai dan natural
- Tidak mengandung SARA
- Tidak promosi berlebihan
- Maksimal 2 paragraf
- Wajib ada teks output
`
}
