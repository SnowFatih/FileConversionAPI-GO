# FileConversionAPI-GO

🚀 **Service de conversion de fichiers en GO**  
Ce projet permet de convertir des fichiers entre différents formats via une **API REST** et une **interface CLI**.

---

## 📌 **Fonctionnalités**

✔️ **Conversion de fichiers texte** (CSV ↔ JSON).  
✔️ **Compression de fichiers (ZIP).**  
✔️ **Conversion d’images** (JPEG ↔ PNG).  
✔️ **API REST et interface CLI**.

### **📌 Technologies utilisées**

- **Bibliothèque `image`** pour manipuler les formats d’images.
- **`encoding/csv` et `encoding/json`** pour les conversions de fichiers texte.
- **Framework `gin-gonic/gin`** pour l’API REST.
- **`cobra`** pour l’interface CLI.

---

# 📌 **Commandes CLI pour tester les conversions**

| **Action**          | **Commande `CLI`**                                                               |
| ------------------- | -------------------------------------------------------------------------------- |
| **CSV → JSON**      | `./file-converter-cli convert fichier-exemple/ip-feuille.csv ip-feuille.json`    |
| **JSON → CSV**      | `./file-converter-cli convert fichier-exemple/ip-feuille.json ip-feuille.csv`    |
| **PNG → JPEG**      | `./file-converter-cli convert fichier-exemple/photo-png.png photo-converted.jpg` |
| **JPEG → PNG**      | `./file-converter-cli convert fichier-exemple/photo-jpg.jpg photo-converted.png` |
| **Compression ZIP** | `./file-converter-cli convert fichier-exemple/ip-feuille.txt ip-feuille.zip`     |

---

# 📌 **Commandes `curl` pour tester les conversions**

ℹ️ **Avant d’exécuter ces commandes, assure-toi que tu es dans le dossier `fichier-exemple` si tes fichiers y sont stockés.**

| **Action**          | **Commande `curl`**                                                                                     |
| ------------------- | ------------------------------------------------------------------------------------------------------- |
| **CSV → JSON**      | `curl -X POST http://localhost:8080/convert/text -F "file=@ip-feuille.csv" -o ip-feuille-exported.json` |
| **JSON → CSV**      | `curl -X POST http://localhost:8080/convert/text -F "file=@ip-feuille.json" -o ip-feuille-exported.csv` |
| **PNG → JPEG**      | `curl -X POST http://localhost:8080/convert/image -F "file=@photo-png.png" -o photo-exported.jpg`       |
| **JPEG → PNG**      | `curl -X POST http://localhost:8080/convert/image -F "file=@photo-jpg.jpg" -o photo-exported.png`       |
| **Compression ZIP** | `curl -X POST http://localhost:8080/compress -F "file=@ip-feuille.txt" -o ip-feuille-compressed.zip`    |

---

# 📌 **Fichier de test Bruno**

Bruno est mon logiciel préféré pour les tests d'API, c'est un jumeau de Postman

Si vous avez l'application Bruno sur votre PC, vous pourrez directement importer le dossier "FileConversion-GO" à la racine du projet pour avoir tous les tests d'endpoints sur votre Bruno

---

## 🚀 **Démarrer l’API REST (`curl.go`)**

Avant d’exécuter des commandes `curl`, il faut démarrer l’API REST.

1️⃣ **Compiler et exécuter l’API REST** :

```sh
go build -o file-converter-api curl.go
./file-converter-api
```

2️⃣ **Tester avec curl** :

```sh
curl -X POST http://localhost:8080/convert/image -F "file=@photo-png.png" -o photo-exported.jpg
```

---

## 🚀 **🚀 Utiliser la CLI (main.go)**

Avant d’exécuter des commandes CLI, compile le programme.

1️⃣ **Compiler la CLI** :

```sh
go build -o file-converter-cli main.go
```

2️⃣ **Tester avec la CLI** :

```sh
./file-converter-cli convert fichier-exemple/photo-png.png photo-converted.jpg
```

Avant d’exécuter des commandes `curl`, il faut démarrer l’API REST.

## 🎯 Résumé

✅ API REST (curl.go) → Tester avec curl après avoir démarré le serveur.
✅ Interface CLI (main.go) → Convertir des fichiers sans serveur.
✅ Deux exécutables distincts (file-converter-api pour l’API et file-converter-cli pour la CLI).
