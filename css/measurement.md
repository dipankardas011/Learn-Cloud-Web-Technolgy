# Understand Absolute versus Relative Units

The last several challenges all set an element's margin or padding with pixels (px). Pixels are a type of length unit, which is what tells the browser how to size or space an item. In addition to px, CSS has a number of different length unit options that you can use.

The two main types of length units are absolute and relative. Absolute units tie to physical units of length. For example, in and mm refer to inches and millimeters, respectively. Absolute length units approximate the actual measurement on a screen, but there are some differences depending on a screen's resolution.

Relative units, such as em or rem, are relative to another length value. For example, em is based on the size of an element's font. If you use it to set the font-size property itself, it's relative to the parent's font-size.

Note: There are several relative unit options that are tied to the size of the viewport. They are covered in the Responsive Web Design Principles section.

# The override styles in CSS
Our pink-text class overrode our body element's CSS declaration!

We just proved that our classes will override the body element's CSS. So the next logical question is, what can we do to override our pink-text class?

Create an additional CSS class called blue-text that gives an element the color blue. Make sure it's below your pink-text class declaration.

Apply the blue-text class to your h1 element in addition to your pink-text class, and let's see which one wins.

Applying multiple class attributes to a HTML element is done with a space between them like this:

class="class1 class2"
Note: It doesn't matter which order the classes are listed in the HTML element.
```CSS
<style>
  body {
    background-color: black;
    font-family: monospace;
    color: green;
  }
  .pink-text {
    color: pink;
  }
  .blue-text {
    color: blue;
  }
</style>
<h1 class="pink-text blue-text">Hello World!</h1>
```


However, the order of the class declarations in the <style> section is what is important. The second declaration will always take precedence over the first. Because .blue-text is declared second, it overrides the attributes of .pink-text



# Improve Compatibility with Browser Fallbacks


```css
<style>
  :root {
    --red-color: red;
  }
  .red-box {
    background: red;
    background: var(--red-color);
    height: 200px;
    width:200px;
  }
</style>
<div class="red-box"></div>
```

# Inherit CSS Variables
When you create a variable, it is available for you to use inside the selector in which you create it. It also is available in any of that selector's descendants. This happens because CSS variables are inherited, just like ordinary properties.

To make use of inheritance, CSS variables are often defined in the :root element.

:root is a pseudo-class selector that matches the root element of the document, usually the html element. By creating your variables in :root, they will be available globally and can be accessed from any other selector in the style sheet.

```css
:root {
    /* Only change code below this line */
    --penguin-belly: pink;
    /* Only change code above this line */
  }
```

then to override in child
```css
  .penguin {
    /* Only change code below this line */
    --penguin-belly: white;
    /* Only change code above this line */
```