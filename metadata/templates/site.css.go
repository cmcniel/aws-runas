package templates

const SiteCss = `
body {
    font-family: Tahoma, Geneva, Verdana, sans-serif;
    margin: 0;
    background-color: #000080;
}

main {
    margin-left: auto;
    margin-right: auto;
    min-width: 30em;
    max-width: 40em;
    padding: 1em 2em;
}

label {
    font-weight: bold;
    display: inline-block;
    width: 7.25em;
}

#advanced-form {
    display: none;
    overflow: hidden;
}

.short-label {
    width: 3em;
}

.w3-input {
    display: inline-block;
    width: 80%;
    margin-bottom: 0.33em;
}

.accordion {
    cursor: pointer;
    padding: 0;
    color: #777;
}

.accordion:after {
    content: '\02795'; /* Unicode character for "plus" sign (+) */
    font-size: 18px;
    color: #777;
    float: left;
    margin-right: 0.33em;
}

.active:after {
    content: "\2796"; /* Unicode character for "minus" sign (-) */
}

#cred-prompt {
    min-width: 22.5em;
    max-width: 27.5em;
    display: block;
}

.cred-input {
    width: 100%;
}

.cred-submit {
    margin-top: 0.5em;
    display: block;
    width: 100%;
    font-size: large;
    font-weight: bold;
    padding: 0.5em 1em;
    border-radius: 0.33em;
    color: white;
    background-color: darkblue;
    border: 1px solid darkblue;
}

.cred-submit:hover {
    background-color: white;
    border: 1px solid darkblue;
    color: darkblue;
}

#mfa-prompt {
    min-width: 14em;
    max-width: 16em;
}

#profile-prompt {
	min-width: 18em;
	max-width: 20em;
}

#alerts {
	position: absolute;
	width: 40%;
	z-index: 9000;
	display: none;
}
`
