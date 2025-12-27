
package converting

import . "TypeGo/core"



func Process(convertData *ConvertData, blockData *BlockData, nestCount int) {
	convertData.AppendString("if err != nil {{"); 
	convertData.NewLineWithTabs(); 
	convertData.AppendString("return "); 
	
	var tokenList []Token  = blockData.Tokens; 
	if len(tokenList) != 0 {
		PrintTokens(convertData, blockData, nestCount); 
		
	}
	convertData.NewLineWithTabs(); 
	convertData.AppendChar('}'); 
	
}
