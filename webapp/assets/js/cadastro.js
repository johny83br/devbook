$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario() {

  event.preventDefault();

  console.log('Criando usuário...');

  var nome = $('#nome').val();
  var email = $('#email').val();
  var nick = $('#nick').val();
  var senha = $('#senha').val();
  var confirmarSenha = $('#confirmar_senha').val();

  if (senha != confirmarSenha) {
    alert('As senhas não conferem!');
    return;
  }

  $.ajax({
    url: '/criar-usuario',
    method: 'POST',
    data: {
      nome: nome,
      email: email,
      nick: nick,
      senha: senha,
      confirmar_senha: confirmarSenha
    },
    success: function () {
      alert('Usuário criado com sucesso!');
      window.location.href = '/login';
    },
    error: function (jqXHR, textStatus, errorThrown) {
      console.error(jqXHR.responseJSON.error);
      alert('Ocorreu um erro ao criar o usuário. Tente novamente.');
    }
  });

}