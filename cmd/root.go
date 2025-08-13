
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/bregydoc/gtranslate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fromLang, toLang string

var rootCmd = &cobra.Command{
	Use:   "translate \"texto a traducir\"",
	Short: "Una CLI sencilla para traducir texto.",
	Long:  `Una herramienta de línea de comandos que traduce texto de un idioma a otro y recuerda los idiomas seleccionados.`,
	// Validamos que el usuario haya proporcionado el texto a traducir
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Texto a traducir es el primer argumento
		textToTranslate := args[0]

		// Comprobar si el usuario especificó nuevos idiomas
		// La función "cmd.Flags().Changed(...)" verifica si la flag fue usada
		if cmd.Flags().Changed("from") {
			viper.Set("fromLang", fromLang) // Guardar el nuevo idioma
		}
		if cmd.Flags().Changed("to") {
			viper.Set("toLang", toLang) // Guardar el nuevo idioma
		}

		// Guardar la configuración si se cambiaron los idiomas
		if cmd.Flags().Changed("from") || cmd.Flags().Changed("to") {
			if err := viper.WriteConfig(); err != nil {
				// Si el archivo de configuración no existe, lo creamos
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					viper.SafeWriteConfig()
				} else {
					fmt.Println("Error al guardar la configuración:", err)
				}
			}
			fmt.Println("Idiomas por defecto actualizados.")
		}

		// Obtener los idiomas de la configuración de Viper
		finalFromLang := viper.GetString("fromLang")
		finalToLang := viper.GetString("toLang")

		// Validar que los idiomas estén configurados
		if finalFromLang == "" || finalToLang == "" {
			fmt.Println("Por favor, configure los idiomas de origen y destino con las flags --from y --to.")
			fmt.Println("Ejemplo: translate --from en --to es \"hello world\"")
			return
		}

		// Realizar la traducción
		fmt.Printf("Traduciendo de '%s' a '%s'...\n", finalFromLang, finalToLang)
		translated, err := gtranslate.TranslateWithParams(
			textToTranslate,
			gtranslate.TranslationParams{
				From: finalFromLang,
				To:   finalToLang,
			},
		)
		if err != nil {
			fmt.Println("Error en la traducción:", err)
			return
		}

		// Imprimir el resultado
		fmt.Println("-----------")
		fmt.Println(translated)
		fmt.Println("-----------")
	},
}


// Execute agrega todos los comandos hijos al comando raíz y establece las flags apropiadamente.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init() se llama antes de main. Aquí configuramos Cobra y Viper.
func init() {
	cobra.OnInitialize(initConfig)

	// Definimos las flags persistentes que estarán disponibles para este comando y todos sus hijos
	rootCmd.PersistentFlags().StringVar(&fromLang, "from", "", "Idioma de origen (ej: en, es, fr)")
	rootCmd.PersistentFlags().StringVar(&toLang, "to", "", "Idioma de destino (ej: en, es, fr)")
}

// initConfig lee el archivo de configuración y las variables de entorno si existen.
func initConfig() {
	// Nombre del archivo de configuración (sin extensión)
	viper.SetConfigName("config")
	// Extensión del archivo
	viper.SetConfigType("yaml")
	// Ruta donde buscar el archivo de configuración
	viper.AddConfigPath("$HOME/.go-translator-cli") // Directorio home del usuario

	// Leer la configuración
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// El archivo de configuración no se encontró; no es un error, se creará después.
		} else {
			// Error al leer el archivo de configuración
			fmt.Println("No se pudo leer el archivo de configuración:", err)
		}
	}
}
