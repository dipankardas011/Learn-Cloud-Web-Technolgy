# Fonts
Name |	Generic type |	Notes
--|--|---
Arial |	sans-serif |	It's often considered best practice to also add Helvetica as a preferred alternative to Arial as, although their font faces are almost identical, Helvetica is considered to have a nicer shape, even if Arial is more broadly available.
Courier New |	monospace |	Some OSes have an alternative (possibly older) version of the Courier New font called Courier. It's considered best practice to use both with Courier New as the preferred alternative.
Georgia |	serif |	
Times New Roman |	serif |	Some OSes have an alternative (possibly older) version of the Times New Roman font called Times. It's considered best practice to use both with Times New Roman as the preferred alternative.
Trebuchet MS |	sans-serif |	You should be careful with using this font — it isn't widely available on mobile OSes.
Verdana |	sans-serif |	

# Finding fonts

For this example, we'll use two web fonts: one for the headings and one for the body text. To start with, we need to find the font files that contain the fonts. Fonts are created by font foundries and are stored in different file formats. There are generally three types of sites where you can obtain fonts:

    A free font distributor: This is a site that makes free fonts available for download (there may still be some license conditions, such as crediting the font creator). Examples include Font Squirrel, dafont, and Everything Fonts.
    A paid font distributor: This is a site that makes fonts available for a charge, such as fonts.com or myfonts.com. You can also buy fonts directly from font foundries, for example Linotype, Monotype, or Exljbris.
    An online font service: This is a site that stores and serves the fonts for you, making the whole process easier. See the Using an online font service section for more details.

Let's find some fonts! Go to Font Squirrel and choose two fonts: a nice interesting font for the headings (maybe a nice display or slab serif font), and a slightly less flashy and more readable font for the paragraphs. When you've found a font, press the download button and save the file inside the same directory as the HTML and CSS files you saved earlier. It doesn't matter whether they are TTF (True Type Fonts) or OTF (Open Type Fonts).

Unzip the two font packages (Web fonts are usually distributed in ZIP files containing the font file(s) and licensing information). You may find multiple font files in the package — some fonts are distributed as a family with different variants available, for example thin, medium, bold, italic, thin italic, etc. For this example, we just want you to concern yourself with a single font file for each choice.

Note: In Font Squirrel, under the "Find fonts" section in the right-hand column, you can click on the different tags and classifications to filter the displayed choices.
# Generating the required code

Now you'll need to generate the required code (and font formats). For each font, follow these steps:

* Make sure you have satisfied any licensing requirement if you are going to use this in a commercial and/or Web project.
* Go to the Fontsquirrel 
* [Webfont Generator](https://www.fontsquirrel.com/tools/webfont-generator)
* Upload your two font files using the Upload Fonts button.
* Check the checkbox labeled "Yes, the fonts I'm uploading are legally eligible for web embedding."
* Click Download your kit.

After the generator has finished processing, you should get a ZIP file to download. Save it in the same directory as your HTML and CSS.

If you need to support legacy browsers, select the "Expert" mode in the Fontsquirrel Webfont Generator, select SVG, EOT, and TTF formats before downloading your kit.

Web services for font generation typically limit file sizes. In such a case, consider using tools such as:

1. [sfnt2woff-zopfli](https://github.com/bramstein/sfnt2woff-zopfli) for converting ttf to woff
1. [fontforge](https://fontforge.org/) for converting from ttf to svg
1. [batik ttf2svffor](https://people.apache.org/~clay/batik/ttf2svg.html) converting from ttf to svg
1. [woff2for](https://github.com/google/woff2) converting from ttf to woff2
