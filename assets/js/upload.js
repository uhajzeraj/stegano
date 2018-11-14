// JavaScript Document

function noPreview() {
  $('#image-preview-div').css("display", "none");
  $('#preview-img').attr('src', 'noimage');
  $('#upload-button').attr('disabled', '');

  $("#chooseOption").fadeOut(300);
  $("#content-wrapper #encode").fadeOut(300);
  $("#content-wrapper #decode").fadeOut(300);
}

function selectImage(e) {
  $('#image-preview-div').css("display", "block");
  $('#preview-img').attr('src', e.target.result);
}


$(document).ready(function () {

  var maxsize = 10 * 1024 * 1024; // 10 MB

    $(`#upload-button`).click(function(e) {

      e.preventDefault();

      var formData = new FormData();
      var text = $(`textarea#text`).val();
      formData.append('text', text);
      var image = $(`#file`)[0].files[0];
      formData.append('image', image);

      $('#message').empty();

      $.ajax({
          url: "stegano",
          type: "POST",
          data: formData,
          processData: false,
          contentType: false,
          success: function(data) {
            if(data == 1) {
              window.location.replace("saved");
            } else {

              var output = "";
              var jsonData = JSON.parse(data);

              // Show all the errors
              for(var i = 0; i < jsonData.length; i++) {
                output += jsonData[i] + "<br/>";
              }
              
              $('#message').html('<div class="alert alert-danger" role="alert">'+output+'</div>');
            }
          }
      });

    });

    // Encode button
    $("#encodeStegano").click(function() {
      event.preventDefault();

      // Disable this button and enable the other
      $(this).prop("disabled", true);
      $("#decodeStegano").prop("disabled", false);

      // Show encode content
      $("#content-wrapper #decode").fadeOut(300, function() {
        $("#content-wrapper #encode").fadeIn(300);
      });
    });

    // Decode button
    $("#decodeStegano").click(function() {
      event.preventDefault();

      $('#message').empty();

      // Disable this button and enable the other
      $(this).prop("disabled", true);
      $("#encodeStegano").prop("disabled", false);

      // Show Decode content
      $("#content-wrapper #encode").fadeOut(300, function() {
        $("#content-wrapper #decode").fadeIn(300);
      });

      // Prepare the image
      var formData = new FormData();
      var image = $(`#file`)[0].files[0];
      formData.append('image', image);

      // Ajax request
      $.ajax({
          url: "steganoDecode",
          type: "POST",
          data: formData,
          processData: false,
          contentType: false,
          success: function(response) {
            $(`textarea#decodedText`).val(response);
          }
      });

    });

    // When an image is selected
    $('#file').change(function() {

      var file = this.files[0];
      var match = ["image/jpeg", "image/png", "image/jpg"];

      if (!((file.type == match[0]) || (file.type == match[1]) || (file.type == match[2]))) {
        noPreview();
        $('#message').html('<div class="alert alert-danger" role="alert">Invalid image format. Allowed formats: JPG, JPEG, PNG.</div>');
        return;
      }

      if (file.size > maxsize) {
        noPreview();
        $('#message').html('<div class=\"alert alert-danger\" role=\"alert\">The size of image you are attempting to upload is ' + (file.size/1024).toFixed(2) + ' KB, maximum size allowed is ' + (maxsize/1024).toFixed(2) + ' KB</div>');
        return;
      }

      // Show the options
      $("#chooseOption").fadeIn(300);

      // In case an image was chosen before, remove previous configs
      $("#content-wrapper #encode").fadeOut(300);
      $("#content-wrapper #decode").fadeOut(300);
      $(`textarea#decodedText`).val("");
      $("#encodeStegano").prop("disabled", false);
      $("#decodeStegano").prop("disabled", false);
 
      $('#message').empty();

      $('#upload-button').removeAttr("disabled");

      var reader = new FileReader();
      reader.onload = selectImage;
      reader.readAsDataURL(this.files[0]);

    });

});

