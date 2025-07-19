$('#nova-publicacao').on('submit', criarPublicacao) ;

$(document).on('click','.curtir-publicacao', curtirPublicacao);
$(document).on('click','.descurtir-publicacao', descurtirPublicacao);

$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

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
    Swal.fire({
      title: 'Erro',
      text: 'Não foi possível criar a publicação. Tente novamente mais tarde.',
      icon: 'error'
    });
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
    Swal.fire("Ops...", "Erro ao curtir a publicação!", "warning");
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
    Swal.fire("Ops...", "Erro ao descurtir a publicação!", "warning");
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
    Swal.fire({
      title: 'Publicação atualizada!',
      text: 'A publicação foi atualizada com sucesso.', 
      icon: 'success'
    }).then(function() {
    window.location = "/home";
    });   
  }).fail(function() {
    Swal.fire("Ops..", "Erro ao atualizar a publicação!", "warning");  
}).always(function() {
    $('#atualizar-publicacao').prop('disabled', false);
  });
}

function deletarPublicacao(evento) {
   evento.preventDefault();

   Swal.fire({
    title: 'Atenção!',
    text: "Deseja apagar a publicação? Essa ação é irreversível!",
    showCancelButton: true,    
    cancelButtonText: 'Cancelar',
    // showConfirmButton: true,
    confirmButtonText: 'Sim, apagar',
    icon: 'warning'
  }).then(function(confirmacao) {
    if (!confirmacao.value) return;

    const elementoClicado = $(evento.target);
    const publicacao = elementoClicado.closest('div');
    const publicacaoId = publicacao.data('publicacao-id');

    elementoClicado.prop('disabled', true);
    $.ajax({
      url: `/publicacoes/${publicacaoId}`,
      method: "DELETE"
    }).done(function(){
      publicacao.fadeOut("slow", function(){
        $(this).remove();
      }); 
    }).fail(function() {
      Swal.fire("Ops.!", "Erro ao excluir a publicação!", "warning");
    });
  })
}
