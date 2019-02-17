This image reader program looks through the images folder in the local directory.
It then asks for which color you are looking for.
Currently supports the 3 basic RGB values.

The program searches for images that contain the color requested.
It then sorts the images by the desired RGB value and outputs the sorted images' names.

Valid images must have extension .png or .jpg.
Supports up to 4k image resolution.

TODO:
Improve runtime
Group image reader into a package
Implement more colors
Output images sorted as requested and then the remaining images sorted by the remaining colors

  Ex.
    Images with highest occurring color and value of that color:
      Green 85
      Red   180
      Blue  15
      Green 52
      Red   95,
      Blue  150
      Red   250
      Blue  92
      Green 10

    Has output:
        Red   250
        Red   180
        Red   95
        Blue  150
        Blue  92
        Blue  15
        Green 85
        Green 52
        Green 10
