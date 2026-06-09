$('#nova-publicacao').on('submit', criarPublicacao)

function criarPublicacao(evento) {

  evento.preventDefault();

  $.ajax({
    url: "/publicacoes",
    method: "POST",
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val()
    },
    success: function () {
      alert('Publicacação criada com sucesso!');
      window.location = "/home";
    },
    error: function (jqXHR, textStatus, errorThrown) {
      alert('Erro ao criar a publicação!');
    }
  });

}