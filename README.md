# Weather CLI

A command-line interface for fetching weather information for various locations.

## Installation

To install the Weather CLI, follow these steps:

1. Ensure you have Go installed on your system.
2. Run the following command to download the necessary modules:
   ```
   go mod download ./...
   ```
3. Build the Weather CLI using the following command:
   ```
   go build ./...
   ```
4. The executable will be generated in your current directory.

## Usage

To use the Weather CLI, simply run the executable followed by the location you're interested in. For example:
   ```
   ./weather-cli London
   ./weather-cli Goa
   ./weather-cli Amsterdam
   ```
Replace `London`, `Amsterdam`, and `Goa` with your desired location.

## API Key

To use the Weather CLI, you will need to obtain an API key from [weatherapi.com](https://www.weatherapi.com). Once you have your API key, create a new file named `.env` in the root directory of the Weather CLI and add the following line:
   ```
   API_KEY=your_api_key
   ```
Replace `your_api_key` with your actual API key.

## Contributing

Contributions are welcome! If you'd like to contribute to the Weather CLI, please fork this repository, make your changes, and submit a pull request.

## License

The Weather CLI is licensed under the MIT License.