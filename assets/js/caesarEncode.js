$(document).ready(function() {
    $(`#upload-button`).click(function(e) {

        e.preventDefault();

        var formData = new FormData();
        var plaintext = $(`textarea#plaintext`).val();
        formData.append('plaintext', plaintext);
        var shiftSize = $(`#shiftSize`).val();
        formData.append('shiftSize', shiftSize);

        $('#loading').show();

        $.ajax({
            url: "caesar",
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