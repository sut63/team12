/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntClubapplication,
    EntClubapplicationFromJSON,
    EntClubapplicationFromJSONTyped,
    EntClubapplicationToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntClubappStatusEdges
 */
export interface EntClubappStatusEdges {
    /**
     * Clubapplication holds the value of the clubapplication edge.
     * @type {Array<EntClubapplication>}
     * @memberof EntClubappStatusEdges
     */
    clubapplication?: Array<EntClubapplication>;
}

export function EntClubappStatusEdgesFromJSON(json: any): EntClubappStatusEdges {
    return EntClubappStatusEdgesFromJSONTyped(json, false);
}

export function EntClubappStatusEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntClubappStatusEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'clubapplication': !exists(json, 'clubapplication') ? undefined : ((json['clubapplication'] as Array<any>).map(EntClubapplicationFromJSON)),
    };
}

export function EntClubappStatusEdgesToJSON(value?: EntClubappStatusEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'clubapplication': value.clubapplication === undefined ? undefined : ((value.clubapplication as Array<any>).map(EntClubapplicationToJSON)),
    };
}

