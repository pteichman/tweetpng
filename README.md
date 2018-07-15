tweetpng
========

This program loads a PNG file, adds some invisible transparency and
strips its metadata, and saves the result as another PNG.

Why? Conventional wisdom is that Twitter will leave a PNG as a PNG
(not convert to PNG) if a file has any transparency. This allow you
to post screenshots without JPEG artifacts.

Install & use
-------------

    $ go get github.com/pteichman/tweetpng
    $ tweetpng /path/to/file.png
