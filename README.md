# background-changer
Download random background image from [unsplash](https://unsplash.com/).

## Prerequisite
1. Create a developer profile on [unsplash](https://unsplash.com/) in order to get an api key.
1. Create a folder where your downloaded image will be stored.
1. In the system preferences, chose to use as background images a custom location and add the folder you have just created, where your image will be stored.
1. Set the `change picture` parameters to preferred (e.g. daily)
1. Create a `.env` file in the project directory and put in it `key=<unsplash_developer_key>`

## Run
Run the script with `go run main.go "<location_to_download_img>"` to download a new image.

In order to automatically run the script, add the execution of it in the crontab.

## Customization
The script downloads images from the collection [nature](https://unsplash.com/collections/880012/nature), modifying the `collectionIDs` variable in the script change this behaviour.