// JavaScript Document

function noPreview() {
  $('#image-preview-div').css("display", "none");
  $('#preview-img').attr('src', 'noimage');
  $('upload-button').attr('disabled', '');
}

function selectImage(e) {
  $('#file').css("color", "green");
  $('#image-preview-div').css("display", "block");
  $('#preview-img').attr('src', e.target.result);
  $('#preview-img').css('max-width', '550px');
}


$(document).ready(function () {

  var maxsize = 5000 * 1024;

  $('#max-size').html((maxsize/1024).toFixed(2));

    $(`#upload-button`).click(function(e) {

        e.preventDefault();

        var formData = new FormData();
        var text = $(`textarea#text`).val();
        formData.append('text', text);
        var image = $(`#file`)[0].files[0];
        formData.append('image', image);

        $('#message').empty();
        $('#loading').show();

        $.ajax({
            url: "stegano",
            type: "POST",
            data: formData,
            processData: false,
            contentType: false,
            success: function(data) {
              
                if(data == 1) {
                  $('#loading').hide();
                  window.location.replace("saved");
                }

            }
        });

    });

  $('#file').change(function() {

    $('#message').empty();

    var file = this.files[0];
    var match = ["image/jpeg", "image/png", "image/jpg"];

    if ( !( (file.type == match[0]) || (file.type == match[1]) || (file.type == match[2]) ) )
    {
      noPreview();

      $('#message').html('<div class="alert alert-warning" role="alert">Unvalid image format. Allowed formats: JPG, JPEG, PNG.</div>');

      return false;
    }

    if ( file.size > maxsize )
    {
      noPreview();

      $('#message').html('<div class=\"alert alert-danger\" role=\"alert\">The size of image you are attempting to upload is ' + (file.size/1024).toFixed(2) + ' KB, maximum size allowed is ' + (maxsize/1024).toFixed(2) + ' KB</div>');

      return false;
    }

    $('#upload-button').removeAttr("disabled");

    var reader = new FileReader();
    reader.onload = selectImage;
    reader.readAsDataURL(this.files[0]);

  });


  // Decode the encrypted file with stegano:
  $('#decryptStegano').click(function (e) {

      e.preventDefault();

      var formData = new FormData();
      var image = $(`#file`)[0].files[0];
      formData.append('image', image);

      $('#encoding').hide();
      $('#decoding').show();

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

  })

    // Decode the encrypted file with stegano:
    $('#returnToEncryption').click(function (e) {

        e.preventDefault();

        $('#decoding').hide();
        $('#encoding').show();
    })

});

