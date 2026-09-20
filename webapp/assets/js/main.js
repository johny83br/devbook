$(document).ready(function () {
  // Pega o pathname atual (ex: "/home", "/perfil")
  var currentPath = window.location.pathname;

  // Normaliza: remove barra final, exceto quando for apenas "/"
  if (currentPath.length > 1 && currentPath.endsWith('/')) {
    currentPath = currentPath.slice(0, -1);
  }

  // Percorre todos os .nav-item com link interno
  $('.nav-item').each(function () {
    var $navItem = $(this);
    var $link = $navItem.find('a.nav-link').first();

    // Ignora itens sem link ou links externos/âncoras
    if (!$link.length) return;

    var href = $link.attr('href');
    if (!href || href.startsWith('http') || href.startsWith('#') || href.startsWith('mailto:')) {
      return;
    }

    // Remove barra final do href para comparação
    var linkPath = href;
    if (linkPath.length > 1 && linkPath.endsWith('/')) {
      linkPath = linkPath.slice(0, -1);
    }

    // Compara caminho exato OU se a rota atual começa com o link
    // (útil para rotas filhas, ex: /perfil/editar marca /perfil)
    if (currentPath === linkPath || currentPath.startsWith(linkPath + '/')) {
      $navItem.addClass('active');
      $link.addClass('active').attr('aria-current', 'page');
    }
  });
});