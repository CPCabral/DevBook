$('#nova-publicacao').on('submit', criarPublicacao) 

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