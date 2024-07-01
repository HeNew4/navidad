# Navidad CLI

¿Alguna vez te has levantado y pensado: "Ojalá tuviera una herramienta en la terminal que me diga si es Navidad" o "Ojalá tuviera una herramienta en la terminal que me diga cuántas semanas faltan para Navidad"? ... ¿No? Bueno, no importa, ¡les presento este hermoso repositorio!

## Descripción

Navidad CLI es una aplicación divertida que te permite verificar si es Navidad hoy y también calcular cuántas semanas faltan para Navidad. Porque, seamos honestos, ¿quién no quiere saber eso?

## Instalación

Para instalar Navidad CLI, sigue estos sencillos pasos:

1. Clona el repositorio:

   ```bash
   git clone https://github.com/HeNew4/navidad.git
   cd navidad
   ```

2. Instala las dependencias:

   ```bash
   go mod tidy
   ```

3. Compila el proyecto:

   ```bash
   go build .
   ```

## 🛠️ Uso

¡Es muy fácil de usar! Solo tienes que ejecutar el binario que acabas de compilar. Aquí tienes algunos ejemplos:

### Verificar si es Navidad

```bash
./navidad es-navidad
```

### Calcular las semanas restantes hasta Navidad

```bash
./navidad semanas-para-navidad
```

## 📁 Archivos Importantes

- **go.mod**: Módulo y dependencias del proyecto.
- **go.sum**: Suma de comprobación de las dependencias.
- **data.go**: Contiene las funciones principales de la aplicación.
- **app.go** y **draw.go**: Otros archivos importantes de la aplicación (asegúrate de revisarlos).

## 🤔 ¿Por qué?

¿Por qué no? ¡La Navidad es una época mágica del año y ahora puedes saber cuánto falta para celebrarla en cualquier momento desde tu terminal!

## 🙌 Contribuciones

¡Las contribuciones son bienvenidas! Si tienes alguna idea divertida o mejoras, no dudes en hacer un fork del repositorio y enviar un pull request.

## 📜 Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.
