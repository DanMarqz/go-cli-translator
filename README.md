# go-cli-translator

## Compile with:

```bash
go build -o {NAME_OUTPUT}
```

## Move to /usr/local/bin:
```bash
sudo mv {NAME_OUTPUT} /usr/local/bin
```

## Create folder for config file:
```bash
mkdir -p ~/.go-translator-cli
```

## Example use:

```bash
{NAME_OUTPUT} --from en --to es "Hello world" 
# Output: 
Idiomas por defecto actualizados.
Traduciendo de 'en' a 'es'...
-----------
Hola Mundo
-----------

# After last command:
{NAME_OUTPUT} "Next phrase to translate"
# Output:
Traduciendo de 'en' a 'es'...
-----------
Siguiente frase para traducir
-----------

# Select new languages:
{NAME_OUTPUT} --from es --to fr "Hola"
# Output:
Idiomas por defecto actualizados.
Traduciendo de 'es' a 'fr'...
-----------
Bonjour
-----------
```
