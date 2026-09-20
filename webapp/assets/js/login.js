$('#login').on('submit', fazerLogin);

function fazerLogin() {

  event.preventDefault();

  console.log('Fazendo login...');

  var email = $('#email').val();
  var senha = $('#senha').val();

  if (email == "") {
    Swal.fire('Erro', 'O e-mail é obrigatório!', 'error');
    return;
  }

  if (senha == "") {
    Swal.fire('Erro', 'A senha é obrigatória!', 'error');
    return;
  }

  $.ajax({
    url: '/login',
    method: 'POST',
    data: {
      email: email,
      senha: senha
    },
    success: function () {
      window.location.href = '/home';
    },
    error: function (jqXHR, textStatus, errorThrown) {
      console.error(jqXHR.responseJSON.error);
      Swal.fire('Erro', 'Ocorreu um erro ao fazer login. Tente novamente.', 'error');
    }
  });

}