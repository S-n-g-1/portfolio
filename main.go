package main

import (
	"os"
)

func main() {
	html := `<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Portofolio - Sugianto</title>
    <style>
        body { font-family: monospace; background: #121212; color: #00ffcc; max-width: 700px; margin: 40px auto; padding: 0 20px; }
        h1 { border-bottom: 1px dashed #00ffcc; padding-bottom: 10px; }
        a { color: #ff007f; text-decoration: none; }
    </style>
</head>
<body>
    <h1>~/sugianto</h1>
    <p><strong>Role:</strong> Software & Game Developer</p>
    <p><strong>Tech Stack:</strong> Go, JavaScript, TypeScript, PHP, C, Python</p>
    
    <h2>// Projects</h2>
    <ul>
        <li><a href="https://github.com/S-n-g-1/ango" target="_blank">ANGO Engine</a> - Custom Go-based visual novel engine</li>
    </ul>

    <h2>// Contact</h2>
    <p>GitHub: <a href="https://github.com/S-n-g-1" target="_blank">@S-n-g-1</a></p>
</body>
</html>`

	err := os.WriteFile("index.html", []byte(html), 0644)
	if err != nil {
		panic(err)
	}
}
