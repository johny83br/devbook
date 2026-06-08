$('#login').on('submit', fazerLogin);

function fazerLogin() {

  event.preventDefault();

  console.log('Fazendo login...');

  var email = $('#email').val();
  var senha = $('#senha').val();

  if (email == "") {
    alert('O e-mail é obrigatório!');
    return;
  }

  if (senha == "") {
    alert('A senha é obrigatória!');
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
      // alert('Login realizado com sucesso!');
      window.location.href = '/home';
    },
    error: function (jqXHR, textStatus, errorThrown) {
      console.error(jqXHR.responseJSON.error);
      alert('Ocorreu um erro ao fazer login. Tente novamente.');
    }
  });

}