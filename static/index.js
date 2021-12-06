console.log('hello go!');
var host = "/day1";
var host2 = "/day2";
$('#submit').click(function () {
    $.ajax({
        url: host + '/login',
        method: 'POST',
        data: {
            username: $('#username').val(),
            password: $('#password').val(),
        },
        success: function (res) {
            console.log(res);
        }
    })
})
$('#submitJson').click(function () {
    $.ajax({
        url: host2 + '/loginJson',
        method: 'POST',
        headers: {
            "Content-Type": "application/json;charset=UTF-8"
        },
        data: JSON.stringify({
            username: $('#username').val(),
            password: $('#password').val(),
        }),
        success: function (res) {
            console.log(res);
        }
    })
})

$('#upload').on('click', function (e) {
    var formData = new FormData();
    var files = $('#file').prop('files');
    console.log(files);
    formData.append('file', files[0])
    $.ajax({
        url: host + '/upload',
        method: 'POST',
        data: formData,
        contentType: false,
        processData: false,
        success: function (res) {
            console.log(res);
        }
    })
})

$('#uploads').click(function (res) {
    var formData = new FormData();
    var files = $('#files').prop('files');
    if (!files || !files.length) return;
    for (var i = 0; i < files.length; i++) {
        formData.append('files[]', files[i])
    }
    $.ajax({
        url: host + '/uploads',
        method: 'POST',
        data: formData,
        contentType: false,
        processData: false,
        success: function (res) {
            console.log(res);
        }
    })
})