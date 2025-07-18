$('#nova-publicacao').on('submit', criarPublicacao) 
$('.curtir-publicacao').on('click', curtirPublicacao)

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
  }).fail(function() {
    alert("Erro ao curtir a publicação!");
  }).always(function() {
    elementoClicado.prop('disabled', false);
  });
}