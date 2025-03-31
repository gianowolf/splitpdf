# splitpdf

`splitpdf` es una herramienta de línea de comandos escrita en Go que permite dividir un archivo PDF en múltiples capítulos definidos por un archivo CSV. Cada capítulo se extrae como un nuevo archivo PDF que contiene el rango de páginas correspondiente.

## 📥 Descargas

Podés descargar los ejecutables desde la [sección de Releases](https://github.com/gianowolf/splitpdf/releases):

- [`splitpdf_windows.zip`](https://github.com/gianowolf/splitpdf/releases) — para sistemas Windows
- [`splitpdf_linux.zip`](https://github.com/gianowolf/splitpdf/releases) — para sistemas Linux

Una vez descargado, descomprimir el archivo y ejecutar el binario desde la terminal.

### Ejemplo en Windows

```powershell
splitpdf_windows.exe -i "documento.pdf" -f "capitulos.csv" -n "Demo"
```

### Ejemplo en Linux 

```bash
chmod +x splitpdf_linux
./splitpdf_linux -i documento.pdf -f capitulos.csv -n Demo
```

## Formato del CSV

El archivo CSV debe tener encabezados y estar separado por comas. Debe contener al menos las siguientes columnas:

```csv
chapter,title,first_page,last_page
01,Introduction,1,10
02,Next Topic,15,24
```

- `chapter`: identificador del capítulo (usado en el nombre del archivo).
- `title`: título del capítulo (usado en el nombre del archivo).
- `first_page`: número de página inicial (comienza en 1).
- `last_page`: número de página final (inclusive).

## Instalación

Se requiere tener [Go](https://golang.org/dl/) instalado.

```bash
git clone https://github.com/gianowolf/splitpdf.git
cd splitpdf
make build
```

También puede compilarse para Windows o Linux con:

```bash
make build-windows
make build-linux
```

## Uso

Desde la línea de comandos:

```bash
./splitpdf -i input.pdf -f chapters.csv [-n prefix]
```

- `-i`: archivo PDF de entrada.
- `-f`: archivo CSV con los capítulos.
- `-n`: (opcional) prefijo para los nombres de los archivos generados.

### Ejemplo

```bash
./splitpdf -i book.pdf -f chapters.csv -n Demo
```

Generará archivos como:

```
Demo_01_Introduction.pdf
Demo_02_NextTopic.pdf
...
```

## Desarrollo

Para compilar y ejecutar localmente:

```bash
make run
```

Esto compilará y ejecutará el binario utilizando el PDF y CSV de prueba definidos en el Makefile.

## Licencia

Este proyecto se distribuye bajo la licencia MIT. Ver `LICENSE` para más información.
