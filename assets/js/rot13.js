$(document).ready(function() {
    $(`#upload-button`).click(function(e) {

        e.preventDefault();

        var formData = new FormData();
        var plaintext = $(`textarea#plaintext`).val();
        formData.append('plaintext', plaintext);

        $('#loading').show();

        $.ajax({
            url: "rot13",
            type: "POST",
            data: formData,
            processData: false,
            contentType: false,
            success: function(response) {
                $('#loading').hide();
                $(`textarea#ciphertext`).val(response);
            }
        });


    });
});