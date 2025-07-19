$('#nova-publicacao').on('submit', criarPublicacao) ;

$(document).on('click','.curtir-publicacao', curtirPublicacao);
$(document).on('click','.descurtir-publicacao', descurtirPublicacao);

$('#atualizar-publicacao').on('click', atualizarPublicacao);

function criarPublicacao(evento) {
  evento.preventDefault();

  $.ajax({
    url: '/publicacoes',
    type: 'POST',
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val()
    }
  }).done(function(){
    window.location.reload();
  }).fail(function() {
    alert('Erro ao criar publicação. Tente novamente mais tarde.');
  })
}

function curtirPublicacao(evento) {
  evento.preventDefault();

  const elementoClicado = $(evento.target);
  const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

  elementoClicado.prop('disabled', true);
  $.ajax({
    url: `/publicacoes/${publicacaoId}/curtir`,
    method: "POST"
  }).done(function(){
    const contadorDeCurtidas = elementoClicado.next('span');
    const numeroDeCurtidas = parseInt(contadorDeCurtidas.text(), 10);

    contadorDeCurtidas.text(numeroDeCurtidas + 1);

    elementoClicado.addClass('text-danger');
    elementoClicado.addClass('descurtir-publicacao');
    elementoClicado.removeClass('curtir-publicacao');
    // elementoClicado.text('Descurtir');

  }).fail(function() {
    alert("Erro ao curtir a publicação!");
  }).always(function() {
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
    method: "POST"
  }).done(function(){
    const contadorDeCurtidas = elementoClicado.next('span');
    const numeroDeCurtidas = parseInt(contadorDeCurtidas.text(), 10);

    contadorDeCurtidas.text(numeroDeCurtidas - 1);

    elementoClicado.removeClass('text-danger');
    elementoClicado.removeClass('descurtir-publicacao');
    elementoClicado.addClass('curtir-publicacao');    

  }).fail(function() {
    alert("Erro ao curtir a publicação!");
  }).always(function() {
    elementoClicado.prop('disabled', false);
  });
}

function atualizarPublicacao(){
  $(this).prop('disabled', true);

  const publicacaoId = $(this).data('publicacao-id');

  $.ajax({
    url: `/publicacoes/${publicacaoId}`,
    method: "PUT",
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val()
    }    
  }).done(function() {
    alert("Publicação atualizada com Sucesso!");    
  }).fail(function() {
    alert('Erro ao atualizar publicação. Tente novamente mais tarde.');
  }).always(function() {
    $('#atualizar-publicacao').prop('disabled', false);
  });
}