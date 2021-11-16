# Flexbox

In Flexbox, flex items will shrink and distribute space between the items according to the space in their container, as their initial behavior. By changing the values for flex-grow and flex-shrink you can indicate how you want the items to behave when they encounter more or less space around them.

# Multicol

The oldest of these layout methods is multicol — when you specify a column-count, this indicates how many columns you want your content to be split into. The browser then works out the size of these, a size that will change according to the screen size.

# CSS grid

In CSS Grid Layout the fr unit allows the distribution of available space across grid tracks. The next example creates a grid container with three tracks sized at 1fr. This will create three column tracks, each taking one part of the available space in the container. You can find out more about this approach to create a grid in the Learn Layout Grids topic, 


So you should always include the above line of HTML in the head of your documents.

There are other settings you can use with the viewport meta tag, however in general the above line is what you will want to use.

> initial-scale: Sets the initial zoom of the page, which we set to 1.

> height: Sets a specific height for the viewport.

> minimum-scale: Sets the minimum zoom level.

> maximum-scale: Sets the maximum zoom level.

> user-scalable: Prevents zooming if set to no.

You should avoid using minimum-scale, maximum-scale, and in particular setting user-scalable to no. Users should be allowed to zoom as much or as little as they need to; preventing this causes accessibility problems.

Media types

The possible types of media you can specify are:

* all
* print
* screen
* speech
```
SYNTAX
@media media-type and (media-feature-rule) {
  /* CSS rules go here */
}
```

The following media query will only set the body to 12pt if the page is printed. It will not apply when the page is loaded in a browser.
```css
@media print {
    body {
        font-size: 12pt;
    }
}
```

The width (and height) media features can be used as ranges, and therefore be prefixed with min- or max- to indicate that the given value is a minimum, or a maximum. For example, to make the color blue if the viewport is narrower than 600 pixels, use max-width:
```css
@media screen and (max-width: 600px) {
    body {
        color: blue;
    }
}
```

One well-supported media feature is orientation, which allows us to test for portrait or landscape mode. To change the body text color if the device is in landscape orientation, use the following media query.
```css
@media (orientation: landscape) {
    body {
        color: rebeccapurple;
    }
}
```

```css
@media (hover: hover) {
    body {
        color: rebeccapurple;
    }
}
```


"and" logic in media queries

To combine media features you can use and in much the same way as we have used and above to combine a media type and feature. For example, we might want to test for a min-width and orientation. The body text will only be blue if the viewport is at least 600 pixels wide and the device is in landscape mode.
```css
@media screen and (min-width: 600px) and (orientation: landscape) {
    body {
        color: blue;
    }
}
```

"or" logic in media queries

If you have a set of queries, any of which could match, then you can comma separate these queries. In the below example the text will be blue if the viewport is at least 600 pixels wide OR the device is in landscape orientation. If either of these things are true the query matches.
```css
@media screen and (min-width: 600px), screen and (orientation: landscape) {
    body {
        color: blue;
    }
}
```

"not" logic in media queries

You can negate an entire media query by using the not operator. This reverses the meaning of the entire media query. Therefore in this next example the text will only be blue if the orientation is portrait.
```css
@media not all and (orientation: landscape) {
    body {
        color: blue;
    }
}
```