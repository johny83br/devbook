$('#nova-publicacao').on('submit', criarPublicacao)
$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);

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

function curtirPublicacao(evento) {

  evento.preventDefault();

  const elementoClicado = $(evento.target);
  const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

  elementoClicado.prop('disabled', true);

  $.ajax({
    url: `/publicacoes/${publicacaoId}/curtir`,
    method: "POST"
  }).done(function () {
    // alert('Publicacação curtida!');

    const contadorCurtidas = elementoClicado.next('span');
    const quantidadeDeCurtidas = parseInt(contadorCurtidas.text());

    contadorCurtidas.text(quantidadeDeCurtidas + 1);

    elementoClicado.addClass('descurtir-publicacao');
    elementoClicado.addClass('text-danger');
    elementoClicado.removeClass('curtir-publicacao');
  }).fail(function () {
    alert('Erro ao curtir a publicação!');
  }).always(function () {
    elementoClicado.prop('disabled', false);
  });

}

function descurtirPublicacao(evento) {

  evento.preventDefault();

  const elementoClicado = $(evento.target);
  const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

  elementoClicado.prop('disabled', true);

  $.ajax({
    url: `/publicacoes/${publicacaoId}/descurtir`,
    method: "POST",
  }).done(function () {
    // alert('Publicacação descurtida!');

    const contadorCurtidas = elementoClicado.next('span');
    const quantidadeDeCurtidas = parseInt(contadorCurtidas.text());

    contadorCurtidas.text(quantidadeDeCurtidas - 1);

    elementoClicado.removeClass('descurtir-publicacao');
    elementoClicado.removeClass('text-danger');
    elementoClicado.addClass('curtir-publicacao');
  }).fail(function () {
    alert('Erro ao descurtir a publicação!');
  }).always(function () {
    elementoClicado.prop('disabled', false);
  });

}