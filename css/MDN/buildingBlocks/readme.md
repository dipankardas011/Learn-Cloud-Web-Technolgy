# Controlling inheritance
CSS provides four special universal property values for controlling inheritance. Every CSS property accepts these values.

## inherit
Sets the property value applied to a selected element to be the same as that of its parent element. Effectively, this "turns on inheritance".

## initial
Sets the property value applied to a selected element to the initial value of that property.

## unset
Resets the property to its natural value, which means that if the property is naturally inherited it acts like inherit, otherwise it acts like initial.

## revert
Acts like unset in many cases, however will revert the property to the browser's default styling rather than the defaults applied to that property.


# Reference table of selectors
The below table gives you an overview of the selectors you have available to use, along with links to the pages in this guide which will show you how to use each type of selector. I have also included a link to the MDN page for each selector where you can check browser support information. You can use this as a reference to come back to when you need to look up selectors later in the material, or as you experiment with CSS generally.

Selector |	Example |	Learn CSS tutorial
--|--|--
Type selector |	h1 {  } |	Type selectors
Universal selector |	* {  } |	The universal selector
Class selector |	.box {  } |	Class selectors
id selector |	#unique { } |	ID selectors
Attribute selector |	a[title] {  } |	Attribute selectors
Pseudo-class selectors |	p:first-child { } |	Pseudo-classes
Pseudo-element selectors |	p::first-line { } |	Pseudo-elements
Descendant combinator |	article p |	Descendant combinator
Child combinator |	article > p |	Child combinator
Adjacent sibling combinator |	h1 + p |	Adjacent sibling
General sibling combinator |	h1 ~ p |	General sibling