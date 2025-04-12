# ⚡ SpeedTest CLI 🚀

> Test your internet speed right from your terminal — built with Go + Cobra 🐹

![Go Version](https://img.shields.io/badge/Go-1.22-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-WIP-yellow)

---

## 📸 Preview

```bash
$ ./speedster download https://speedtest.tele2.net/5MB.zip

Starting download test from: https://speedtest.tele2.net/5MB.zip
Downloaded 5.00 MB in 1.37 seconds
Download speed: 29.20 Mbps
```

---

## 💡 Features (More Coming Soon!)

✅ Download speed test
⬜ Upload speed test
⬜ Ping + Jitter check
⬜ Auto-select fastest server
⬜ JSON output + Logs
⬜ Fancy progress bar

---

## 🔧 Installation

```bash
git clone https://github.com/parthkapoor-dev/speedster.git
cd speedster
go install
```

> Make sure you have Go installed (v1.18+). Get it here: https://go.dev/dl/

---

## 🚀 Usage

### Test Download Speed:
```bash
./speedster download https://speedtest.tele2.net/10MB.zip
```

### Run Help
```bash
./speedster help
```

---

## 🌐 Test File Sources

These are free public test files you can use:
- [Tele2 Speed Test Files](https://speedtest.tele2.net/)
  - `https://speedtest.tele2.net/1MB.zip`
  - `https://speedtest.tele2.net/5MB.zip`
  - `https://speedtest.tele2.net/10MB.zip`

---

## 🧠 How It Works

1. ⏱ Measures time to download file
2. 📦 Calculates size in bytes
3. 💡 Formula:
   ```
   Mbps = (Bytes * 8) / Seconds / 1,000,000
   ```

---

## 🛠 Tech Stack

- [Golang](https://go.dev/) — performance + concurrency
- [Cobra](https://github.com/spf13/cobra) — CLI framework
- [net/http](https://pkg.go.dev/net/http) — for downloading data

---

## 📌 Project Roadmap

| Phase | Features                    |
|-------|-----------------------------|
| ✅ 1  | CLI + download speed test   |
| ⏳ 2  | Upload speed test           |
| ⏳ 3  | Ping & jitter measurement   |
| ⏳ 4  | Smart server selection      |
| ⏳ 5  | JSON/Pretty output          |

---

## 🤝 Contributing

Pull requests are welcome! Ideas, issues, and features are always appreciated.

```bash
git checkout -b feature/cool-idea
git commit -m "Add cool idea"
git push origin feature/cool-idea
```

---

## 📄 License

MIT © [Your Name](https://github.com/yourname)

---

### 🚨 Star this project if you like terminal magic! 🌟
```

---

### Want it as a file?

I can also drop this into a `README.md` file scaffold or paste it as a downloadable file. Would you like that? Or should I tweak the style to match a personal or professional tone?
