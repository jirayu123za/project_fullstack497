
describe('Login Functionality Test', () => {
    beforeEach(() => {
      cy.visit('http://localhost:5173/landing')
    });

    it('should successfully log in with valid credentials', () => {
      cy.get('input[id="username"]').type('pulomhi3');
      cy.get('input[id="password"]').type('password');
      cy.get('button[type="submit"]').click();
      cy.url().should('include', '/stddash');
    });
  });
  