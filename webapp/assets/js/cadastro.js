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
    Swal.fire('Erro', 'As senhas não conferem!', 'error');
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
      Swal.fire('Sucesso', 'Usuário criado com sucesso!', 'success').then(function () {
        window.location.href = '/login';
      });
    },
    error: function (jqXHR, textStatus, errorThrown) {
      console.error(jqXHR.responseJSON.error);
      Swal.fire('Erro', 'Ocorreu um erro ao criar o usuário. Tente novamente.', 'error');
    }
  });

}