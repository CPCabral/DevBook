$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
  evento.preventDefault();

  if ($('#senha').val() != $('#confirmar-senha').val()) {
    Swal.fire("Ops...", "As senhas não coincidem!", "error");
    return
  }
  $.ajax({
    url: "/usuarios",
    method: "POST",
    data: {
      nome: $('#nome').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      senha: $('#senha').val()
    }
  }).done(function() {
    Swal.fire({
      title: "Sucesso!",
      text: "Usuário cadastrado com sucesso!", 
      icon: "success"
    }).then(function() {
      $.ajax({
        url: "login",
        method: "POST",
        data: {
          email: $('#email').val(),
          senha: $('#senha').val()
        }
      }).done(function() {
        window.location = "/home";        
      }).fail(function(erro) {
        Swal.fire("Ops...", "Erro ao autenticar usuário!", "error");
      });
    }); 
  }).fail(function(erro) {
    Swal.fire("Ops...", "Erro ao cadastrar usuário!", "warning");
  });
}