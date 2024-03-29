# AnRegion

If you ever want to re-render just a region of an animation and want a simple tool to combine the new region-rendered image sequence with the original one, look no further!

![anregion](./anregion.png)

## How To

- Download the [latest released zip file](https://github.com/SimonStorlSchulke/AnRegion/releases/latest), unpack it and run anregion.exe
- Fill in the filepaths to the original and the region-rendered image sequence folders
- If you want, you can let AnRegion ignore all filenames containing the words specified in the third input-field (separate words with a single ; and NO space)
- If you want AnRegion to overwrite the original files, check the checkbox. Else, the composited files will be saved in a subfolder of the Original Renderings Folder */anrender*
- This also works, when only some region-rendered frames (like frame 20-50) are given. AnRegion chooses Images to combine based on the number at the end of the filename.

Supported File Types include exr, png, jpg, tiff - every format that ImageMagick supports.

![anregion](./anregion_sh.jpg)

## Credits

This Tool is written in [Go](https://go.dev) and [Svelte](https://svelte.dev), using [Wails](https://wails.io) and [ImageMagick](https://imagemagick.org). See the ImageMagick license at https://imagemagick.org/script/license.php
