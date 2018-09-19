<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:template match="/">
<html> 
<body>
<h2>OpenPositions</h2>
  <table border="1">
    <tr bgcolor="#9acd32">
      <th style="text-align:left">Symbol</th>
    </tr>
    <xsl:for-each select="FlexQueryResponse/FlexStatements/FlexStatement/OpenPositions/OpenPosition">
    <tr>
      <td><xsl:value-of select="@symbol"/></td>
    </tr>
    </xsl:for-each>
  </table>

<h2>Trades</h2>
  <table border="1">
    <tr bgcolor="#9acd32">
      <th style="text-align:left">Symbol</th>
    </tr>
    <xsl:for-each select="FlexQueryResponse/FlexStatements/FlexStatement/Trades/Trade">
    <tr>
      <td><xsl:value-of select="@symbol"/></td>
    </tr>
    </xsl:for-each>
  </table>

<h2>TransactionTaxes</h2>
  <table border="1">
    <tr bgcolor="#9acd32">
      <th style="text-align:left">Symbol</th>
    </tr>
    <xsl:for-each select="FlexQueryResponse/FlexStatements/FlexStatement/TransactionTaxes/TransactionTax">
    <tr>
      <td><xsl:value-of select="@symbol"/></td>
    </tr>
    </xsl:for-each>
  </table>

  <h2>CashTransactions</h2>
  <table border="1">
    <tr bgcolor="#9acd32">
      <th style="text-align:left">Date</th>
      <th style="text-align:left">Amount</th>
    </tr>
    <xsl:for-each select="FlexQueryResponse/FlexStatements/FlexStatement/CashTransactions/CashTransaction">
    <tr>
      <td><xsl:value-of select="@dateTime"/></td>
      <td><xsl:value-of select="@amount"/></td>
    </tr>
    </xsl:for-each>
  </table>
</body>
</html>
</xsl:template>
</xsl:stylesheet>

