import { adminLogin, createProject, logout } from '../support/commands'

describe('global projects filter', () => {
  const proj1  = "cypress-project-1"
  const proj2  = "cypress-project-2"
  const proj3  = "cypress-project-3"
  // TODO uncomment with non-admin test
  // const pol_id = "cypress-policy"
  // const nonAdminUsername = "nonadmin"

  before(() => {
    adminLogin('/').then(() => {
      let admin = JSON.parse(localStorage.getItem('chef-automate-user'))
      cleanupProjects(admin.id_token)
      createProject(admin.id_token, proj1)
      createProject(admin.id_token, proj2)
      createProject(admin.id_token, proj3)
      // TODO uncomment with non-admin test/ move up project creation
      // createUser(admin.id_token, nonAdminUsername)
      // createPolicy(admin.id_token, pol_id, nonAdminUsername, [proj1, proj2])
      logout()
    })
  })

  it('shows all projects for admin', () => {
    adminLogin('/settings')

    cy.get('chef-sidebar')
      .should('have.attr', 'minor-version')
      .then((version) => {
        if (version === 'v1') {
          cy.get('[data-cy=projects-filter-button]').click()
          
          const allowedProjects = [proj1, proj2, proj3, '(unassigned)'];
          // we don't check that projects in dropdown match *exactly* as
          // we can't control creation of other projects in the test env
          allowedProjects.forEach(project => {
            cy.get('[data-cy=projects-filter-dropdown]').contains(project)
          })
        }
      })
    logout()
  })

  // TODO can uncomment when we have a flag to remove legacy policies,
  // which are currently allowing full project access to all users.
  // it('shows allowed projects for non-admin', () => {
  //   login('/settings', nonAdminUsername)
  //   // hide modal unrelated to test flow
  //   cy.get('app-welcome-modal').invoke('hide')

  //   cy.get('chef-sidebar')
  //     .should('have.attr', 'minor-version')
  //     .then((version) => {
  //       if (version === 'v1') {
  //         cy.get('[data-cy=projects-filter-button]').click()

  //         const allowedProjects = [proj1, proj2];
  //         cy.get('[data-cy=projects-filter-dropdown] chef-checkbox')
  //           .should(($elements) => { expect($elements).to.have.length(allowedProjects.length) })
  //         allowedProjects.forEach(project => {
  //           cy.get('[data-cy=projects-filter-dropdown]').contains(project)
  //         })
  //       }
  //     })
  //   logout()
  // })
})
