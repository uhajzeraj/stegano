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

    $('#decryption').click(function (e) {

        e.preventDefault();

        $('#encCaesar').hide();
        $('#decCaesar').show();

        $('#encoding').hide();
        $('#decoding').show();

    })

    // Decode the encrypted file with stegano:
    $('#decryptCaesar').click(function (e) {

        e.preventDefault();

        var formData = new FormData();
        var ciphertext = $(`textarea#ciphertext`).val();
        formData.append('ciphertext', ciphertext);
        var shiftSize = $(`#shiftSizeD`).val();
        formData.append('shiftSizeD', shiftSize);

        $.ajax({
            url: "caesarDecode",
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

        $('#encCaesar').show();
        $('#decCaesar').hide();

        $('#encoding').show();
        $('#decoding').hide();
    })
});