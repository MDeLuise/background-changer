# background-changer
Download random background image from [unsplash](https://unsplash.com/).

## Prerequisite
1. Create a developer profile on [unsplash](https://unsplash.com/) in order to get an api key.
1. Create a folder where your downloaded image will be stored.
1. In the system preferences, chose to use as background images a custom location and add the folder you have just created, where your image will be stored.
1. Set the `change picture` parameters to preferred (e.g. daily)
1. Create a `.env` file in the project directory and put in it `key=<unsplash_developer_key>`

## Run
Run the script with `go run main.go` to download a new image in the current directory with the predefined name `<unix_timestamp>.png`.

In order to automatically run the script, add the execution of it in the crontab.

## Customization
The default behaviour consists in downloading images from the collection [nature](https://unsplash.com/collections/880012/nature) in the current directory. Passing command line arguments to the script change this behaviour:
* `-name=<img_name>` change the name of the downloaded image
* `-directory=<path>` change the target directory used to download the images
* `-collection="<id1, id2, ...>"` filter the downloaded image using the specified collection IDs
* `-clean="<true|false>"` remove old downloaded images