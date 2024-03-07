function HtmlEntityDeCoder(text){
    //Create textarea
    const _textArea = document.createElement('textarea')
    //Put Html Entity to InnerHtml it will autometically decoded
    _textArea.innerHTML = text
    //Return textarea's value
    return _textArea.value
}