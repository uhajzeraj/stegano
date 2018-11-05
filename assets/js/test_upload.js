$(document).ready(function() {
  $(`#upload-button`).click(function(e) {
    
    e.preventDefault();

    var formData = new FormData();
    var text = $(`textarea#text`).val();
    formData.append('text', text);
    var image = $(`#file`)[0].files[0];
    formData.append('image', image);

    $.ajax({
      url: "stegano",
      type: "POST",
      data: formData,
      processData: false,
      contentType: false,
      success: function() {
        // Do something if it goes through
      }
    });
    
  });
});