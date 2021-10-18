# HTML validation

So you can see from the above example that you really want to make sure your HTML is well-formed! But how? In a small example like the one seen above, it is easy to search through the lines and find the errors, but what about a huge, complex HTML document?

The best strategy is to start by running your HTML page through the [Markup Validation Service](https://validator.w3.org/) — created and maintained by the W3C, the organization that looks after the specifications that define HTML, CSS, and other web technologies. This webpage takes an HTML document as an input, goes through it, and gives you a report to tell you what is wrong with your HTML.

The HTML validator homepage

To specify the HTML to validate, you can provide a web address, upload an HTML file, or directly input some HTML code.

## Active learning: Validating an HTML document

Let's try this with our [sample document](https://github.com/mdn/learning-area/blob/master/html/introduction-to-html/debugging-html/debug-example.html)

  - First, load the [Markup Validation Service](https://validator.w3.org/) in one browser tab, if it isn't already open.
  - Switch to the [Validate](https://validator.w3.org/#validate_by_input) by Direct Input tab.
  - Copy all of the sample document's code (not just the body) and paste it into the large text area shown in the Markup Validation Service.
  - Press the Check button.

This should give you a list of errors and other information.
![example](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Debugging_HTML/validation-results.png)
A list of HTML validation results from the W3C markup validation service
Interpreting the error messages

The error messages are usually helpful, but sometimes they are not so helpful; with a bit of practice you can work out how to interpret these to fix your code. Let's go through the error messages and see what they mean. You'll see that each message comes with a line and column number to help you to locate the error easily.

  * "Consider adding a lang attribute to the html start tag to declare the language of this document.": This is not an error but a warning. The [recommendation](https://www.w3.org/International/questions/qa-html-language-declarations) is to always define a language, and this warning explains how to do it.
  * "End tag li implied, but there were open elements" (2 instances): These messages indicate that an element is open that should be closed. The ending tag is implied, but not actually there. The line/column information points to the first line after the line where the closing tag should really be, but this is a good enough clue to see what is wrong.
  * "Unclosed element strong": This is really easy to understand — a < strong> element is unclosed, and the line/column information points right to where it is.
  * "End tag strong violates nesting rules": This points out the incorrectly nested elements, and the line/column information points out where they are.
  * "End of file reached when inside an attribute value. Ignoring tag": This one is rather cryptic; it refers to the fact that there is an attribute value not properly formed somewhere, possibly near the end of the file because the end of the file appears inside the attribute value. The fact that the browser doesn't render the link should give us a good clue as to what element is at fault.
  * "End of file seen and there were open elements": This is a bit ambiguous, but basically refers to the fact there are open elements that need to be properly closed. The line numbers point to the last few lines of the file, and this error message comes with a line of code that points out an example of an open element:

  example: < a href="https://www.mozilla.org/>link to Mozilla homepage</ a> ↩ </ ul>↩ </ body>↩</ html>

  **Note**: An attribute missing a closing quote can result in an open element because the rest of the document is interpreted as the attribute's content.
    "Unclosed element ul": This is not very helpful, as the < ul> element is closed correctly. This error comes up because the < a> element is not closed, due to the missing closing quote mark.

If you can't work out what every error message means, don't worry about it — a good idea is to try fixing a few errors at a time. Then try revalidating your HTML to show what errors are left. Sometimes fixing an earlier error will also get rid of other error messages — several errors can often be caused by a single problem, in a domino effect.

You will know when all your errors are fixed when you see the following banner in your output:

Banner that reads "The document validates according to the specified schema(s) and to additional constraints checked by the validator." 